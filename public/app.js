document.addEventListener('DOMContentLoaded', () => {
    // Initialize Lucide icons
    lucide.createIcons();

    const urlInput = document.getElementById('urlInput');
    const scanBtn = document.getElementById('scanBtn');
    const loadingState = document.getElementById('loading');
    const hudContainer = document.getElementById('results');

    // HUD Elements
    const navList = document.getElementById('nav-list');
    const contentHeader = document.getElementById('content-header');
    const contentIcon = document.getElementById('content-icon');
    const contentTitle = document.getElementById('content-title');
    const contentBody = document.getElementById('content-body');

    // Store current scan data globally for fast tab switching
    let currentScanData = {};

    scanBtn.addEventListener('click', initiateScan);
    urlInput.addEventListener('keypress', (e) => {
        if (e.key === 'Enter') initiateScan();
    });

    async function initiateScan() {
        const url = urlInput.value.trim();

        if (!url) {
            urlInput.parentElement.parentElement.style.borderColor = 'var(--status-red)';
            setTimeout(() => {
                urlInput.parentElement.parentElement.style.borderColor = 'var(--glass-border)';
            }, 1000);
            return;
        }

        // UI Transitions
        scanBtn.disabled = true;
        scanBtn.style.opacity = '0.7';
        hudContainer.classList.remove('visible');
        setTimeout(() => {
            hudContainer.classList.add('hidden');
            navList.innerHTML = '';
            resetContentPane();
        }, 300);

        loadingState.classList.remove('hidden');

        try {
            const response = await fetch(`/api/scan?url=${encodeURIComponent(url)}`);

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const data = await response.json();
            currentScanData = data;
            buildHUD(data);

        } catch (error) {
            console.error('Error during scan:', error);
            alert("Failed to reach scanner backend. Is the Go server running?");
        } finally {
            loadingState.classList.add('hidden');
            scanBtn.disabled = false;
            scanBtn.style.opacity = '1';
        }
    }

    function resetContentPane() {
        contentIcon.setAttribute('data-lucide', 'globe');
        contentIcon.className = 'card-icon blue';
        contentTitle.textContent = 'Waiting for module selection...';
        contentBody.innerHTML = `
            <div class="empty-state">
                <i data-lucide="mouse-pointer-click"></i>
                <p>Select an OSINT module from the sidebar to view detailed intelligence.</p>
            </div>
        `;
        lucide.createIcons();
    }

    // Helper functions to map backend plugin keys to UI aesthetics
    function getModuleMeta(key) {
        const rules = {
            'url': { icon: 'globe', color: 'blue', title: 'Target Overview' },
            'missing_headers': { icon: 'shield-alert', color: 'yellow', title: 'Security Headers' },
            'exposed_files': { icon: 'file-warning', color: 'red', title: 'Sensitive Files' },
            'dns_records': { icon: 'network', color: 'blue', title: 'DNS Records' },
            'whois_info': { icon: 'book', color: 'yellow', title: 'WHOIS Registration' },
            'open_ports': { icon: 'radio-tower', color: 'red', title: 'Open Ports' },
            'tech_stack': { icon: 'cpu', color: 'blue', title: 'Technology Stack' },
            'robots_txt': { icon: 'bot', color: 'yellow', title: 'Robots.txt Paths' }
        };
        return rules[key] || { icon: 'server', color: 'pink', title: formatKeyAsTitle(key) };
    }

    function formatKeyAsTitle(key) {
        return key.split('_').map(w => w.charAt(0).toUpperCase() + w.slice(1)).join(' ');
    }

    function determineStatus(key, val) {
        if (!val || (Array.isArray(val) && val.length === 0) || Object.keys(val).length === 0) {
            return 'green'; // Safe/Blank
        }

        if (val.error) return 'yellow'; // Warning/Timeout

        // Domain-specific threat logic
        if (key === 'exposed_files' || key === 'open_ports') return 'red';
        if (key === 'missing_headers') return 'yellow';

        return 'green'; // General data like DNS or Tech Stack
    }

    function buildHUD(data) {
        let firstTabKey = null;

        // Ensure "url" is always first in the sidebar if it exists
        const keys = Object.keys(data).sort((a, b) => {
            if (a === 'url') return -1;
            if (b === 'url') return 1;
            return 0;
        });

        keys.forEach(key => {
            const val = data[key];
            const meta = getModuleMeta(key);
            const statusColor = determineStatus(key, val);

            if (!firstTabKey) firstTabKey = key;

            // Build Sidebar Nav Item
            const li = document.createElement('li');
            li.className = 'nav-item';
            li.dataset.modkey = key;
            li.innerHTML = `
                <i data-lucide="${meta.icon}" class="nav-icon"></i>
                <span class="nav-title">${meta.title}</span>
                <div class="status-dot status-${statusColor}"></div>
            `;

            li.addEventListener('click', () => {
                // Update active state
                document.querySelectorAll('.nav-item').forEach(el => el.classList.remove('active'));
                li.classList.add('active');

                // Render content pane
                renderContentPane(key, val, meta, statusColor);
            });

            navList.appendChild(li);
        });

        // Show HUD
        hudContainer.classList.remove('hidden');
        setTimeout(() => {
            hudContainer.classList.add('visible');
            hudContainer.scrollIntoView({ behavior: 'smooth', block: 'start' });

            // Auto-click the first tab
            if (firstTabKey) {
                const firstTab = document.querySelector('.nav-item');
                if (firstTab) firstTab.click();
            }
        }, 50);

        lucide.createIcons();
    }

    function renderContentPane(key, data, meta, statusColor) {
        // Force a DOM reflow for the CSS animation to re-trigger
        contentBody.style.animation = 'none';
        contentBody.offsetHeight; /* trigger reflow */
        contentBody.style.animation = null;

        // Update Header
        contentIcon.setAttribute('data-lucide', meta.icon);
        contentIcon.className = `card-icon ${meta.color}`;
        contentTitle.textContent = meta.title;

        // Build Body HTML
        let bodyHTML = '';

        if (data.error) {
            bodyHTML = `
                <div class="badge warning">PLUGIN ERROR / TIMEOUT</div>
                <ul class="item-list">
                    <li class="list-item warning"><i data-lucide="alert-triangle"></i> ${data.error}</li>
                </ul>
            `;
        } else if (statusColor === 'green' && (Array.isArray(data) && data.length === 0 || Object.keys(data).length === 0)) {
            let msg = "No findings to report.";
            if (key === 'missing_headers') msg = "All essential security headers present.";
            if (key === 'exposed_files') msg = "No common sensitive files exposed.";
            if (key === 'open_ports') msg = "All common ports filtered/closed.";

            bodyHTML = `
                <div class="badge">SECURE</div>
                <ul class="item-list">
                    <li class="list-item good"><i data-lucide="check-circle"></i> ${msg}</li>
                </ul>
            `;
        } else {
            if (key === 'url') {
                bodyHTML = `<p class="data-label">URL Scanned:</p><p class="data-value highlight">${data}</p>`;
            } else if (Array.isArray(data)) {
                // Arrays
                bodyHTML += `<ul class="item-list">`;
                data.forEach(item => {
                    let liClass = "list-item";
                    let listIcon = "info";

                    if (statusColor === 'red') { liClass += ' critical'; listIcon = 'flame'; }
                    else if (statusColor === 'yellow') { liClass += ' warning'; listIcon = 'alert-circle'; }

                    bodyHTML += `<li class="${liClass}"><i data-lucide="${listIcon}"></i>${item}</li>`;
                });
                bodyHTML += `</ul>`;
            } else if (typeof data === 'object') {
                // Objects
                bodyHTML += `<ul class="item-list">`;
                for (const [subKey, subVal] of Object.entries(data)) {
                    let fmtVal = subVal;
                    if (Array.isArray(subVal)) fmtVal = subVal.join(', ');
                    bodyHTML += `<li class="list-item"><strong>${subKey}:</strong><span style="margin-left:auto; text-align:right;">${fmtVal}</span></li>`;
                }
                bodyHTML += `</ul>`;
            } else {
                // Strings
                bodyHTML = `<p class="data-value">${data}</p>`;
            }
        }

        // Inject and re-parse icons
        contentBody.innerHTML = bodyHTML;
        lucide.createIcons();
    }
});
