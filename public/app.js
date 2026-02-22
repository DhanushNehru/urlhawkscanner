document.addEventListener('DOMContentLoaded', () => {
    // Initialize Lucide icons
    lucide.createIcons();

    const urlInput = document.getElementById('urlInput');
    const scanBtn = document.getElementById('scanBtn');
    const loadingState = document.getElementById('loading');
    const resultsGrid = document.getElementById('results');

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
        resultsGrid.classList.remove('visible');
        setTimeout(() => {
            resultsGrid.classList.add('hidden');
            resultsGrid.innerHTML = ''; // Clear old results dynamically
        }, 300);

        loadingState.classList.remove('hidden');

        try {
            const response = await fetch(`/api/scan?url=${encodeURIComponent(url)}`);

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const data = await response.json();
            renderResults(data);

        } catch (error) {
            console.error('Error during scan:', error);
            alert("Failed to reach scanner backend. Is the Go server running?");
        } finally {
            loadingState.classList.add('hidden');
            scanBtn.disabled = false;
            scanBtn.style.opacity = '1';
        }
    }

    // Helper functions to map backend plugin keys to UI aesthetics
    function getCardStyles(key) {
        const rules = {
            'url': { icon: 'globe', color: 'blue', title: 'Target Overview' },
            'missing_headers': { icon: 'shield-alert', color: 'yellow', title: 'Missing Security Headers' },
            'exposed_files': { icon: 'file-warning', color: 'red', title: 'Exposed Sensitive Files' },
            'dns_records': { icon: 'network', color: 'blue', title: 'DNS Records' },
            'whois_info': { icon: 'book', color: 'yellow', title: 'WHOIS Registration' },
            'open_ports': { icon: 'radio-tower', color: 'red', title: 'Open Ports' },
            'tech_stack': { icon: 'cpu', color: 'blue', title: 'Technology Stack' },
            'robots_txt': { icon: 'bot', color: 'yellow', title: 'Disallowed Paths (Robots)' }
        };
        return rules[key] || { icon: 'server', color: 'pink', title: formatKeyAsTitle(key) };
    }

    function formatKeyAsTitle(key) {
        return key.split('_').map(w => w.charAt(0).toUpperCase() + w.slice(1)).join(' ');
    }

    function renderResults(data) {
        // ALWAYS render target overview first if present
        if (data.url) {
            renderCard('url', data.url);
        }

        // Loop over the rest of the dynamic plugin keys
        for (const [key, val] of Object.entries(data)) {
            if (key === 'url') continue; // Handled above
            if (!val || (Array.isArray(val) && val.length === 0) || Object.keys(val).length === 0) {
                // Feature found nothing safe/useful to report (or timed out empty)
                renderSafeCard(key);
            } else if (val.error) {
                renderErrorCard(key, val.error);
            } else {
                renderCard(key, val);
            }
        }

        // Re-init icons for newly added HTML
        lucide.createIcons();

        // Show Results
        resultsGrid.classList.remove('hidden');
        setTimeout(() => {
            resultsGrid.classList.add('visible');
            resultsGrid.scrollIntoView({ behavior: 'smooth', block: 'start' });
        }, 50);
    }

    function renderSafeCard(key) {
        const style = getCardStyles(key);
        let msg = "No findings to report.";
        if (key === 'missing_headers') msg = "All essential security headers present.";
        if (key === 'exposed_files') msg = "No common sensitive files exposed.";
        if (key === 'open_ports') msg = "All common ports filtered/closed.";

        const cardHTML = `
            <div class="card glass-panel">
                <div class="card-header">
                    <i data-lucide="${style.icon}" class="card-icon ${style.color}"></i>
                    <h2>${style.title}</h2>
                </div>
                <div class="card-body">
                    <div class="badge">SECURE</div>
                    <ul class="item-list">
                        <li class="list-item good">
                            <i data-lucide="check-circle"></i> ${msg}
                        </li>
                    </ul>
                </div>
            </div>
        `;
        resultsGrid.innerHTML += cardHTML;
    }

    function renderErrorCard(key, errorMsg) {
        const style = getCardStyles(key);
        const cardHTML = `
            <div class="card glass-panel">
                <div class="card-header">
                    <i data-lucide="${style.icon}" class="card-icon ${style.color}"></i>
                    <h2>${style.title}</h2>
                </div>
                <div class="card-body">
                    <div class="badge danger">PLUGIN ERROR / TIMEOUT</div>
                    <ul class="item-list">
                        <li class="list-item warning">
                            <i data-lucide="alert-triangle"></i> ${errorMsg}
                        </li>
                    </ul>
                </div>
            </div>
        `;
        resultsGrid.innerHTML += cardHTML;
    }

    function renderCard(key, data) {
        const style = getCardStyles(key);

        let bodyHTML = '';

        if (key === 'url') {
            bodyHTML = `<p class="data-label">URL Scanned:</p><p class="data-value highlight">${data}</p>`;
        } else if (Array.isArray(data)) {
            // Arrays: Missing Headers, Exposed files, Ports, Tech Stack, Robots...
            bodyHTML += `<ul class="item-list">`;
            data.forEach(item => {
                let liClass = "list-item";
                if (key === 'exposed_files' || key === 'open_ports') liClass += ' critical';
                else if (key === 'missing_headers') liClass += ' warning';

                let icon = 'info';
                if (key === 'missing_headers') icon = 'alert-circle';
                if (key === 'exposed_files' || key === 'open_ports') icon = 'flame';

                bodyHTML += `<li class="${liClass}"><i data-lucide="${icon}"></i>${item}</li>`;
            });
            bodyHTML += `</ul>`;
        } else if (typeof data === 'object') {
            // Objects: DNS Records, WHOIS...
            bodyHTML += `<ul class="item-list">`;
            for (const [subKey, subVal] of Object.entries(data)) {
                let fmtVal = subVal;
                if (Array.isArray(subVal)) fmtVal = subVal.join(', ');
                bodyHTML += `<li class="list-item"><strong>${subKey}:</strong> ${fmtVal}</li>`;
            }
            bodyHTML += `</ul>`;
        } else {
            // Strings
            bodyHTML = `<p class="data-value">${data}</p>`;
        }

        const cardHTML = `
            <div class="card glass-panel">
                <div class="card-header">
                    <i data-lucide="${style.icon}" class="card-icon ${style.color}"></i>
                    <h2>${style.title}</h2>
                </div>
                <div class="card-body">
                    ${bodyHTML}
                </div>
            </div>
        `;
        resultsGrid.innerHTML += cardHTML;
    }
});
