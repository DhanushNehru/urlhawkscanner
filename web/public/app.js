document.addEventListener('DOMContentLoaded', () => {
    // Initialize Lucide icons
    lucide.createIcons();

    const urlInput = document.getElementById('urlInput');
    const scanBtn = document.getElementById('scanBtn');
    const loadingState = document.getElementById('loading');
    const resultsGrid = document.getElementById('results');

    // Result elements
    const resUrl = document.getElementById('res-url');
    const headersContent = document.getElementById('headers-content');
    const filesContent = document.getElementById('files-content');

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
        setTimeout(() => resultsGrid.classList.add('hidden'), 300);

        loadingState.classList.remove('hidden');

        try {
            // API Call to the Go Backend
            // Handled relative to where it was served from
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

    function renderResults(data) {
        // 1. Overview
        resUrl.textContent = data.url || 'Unknown';

        // 2. Headers
        headersContent.innerHTML = '';
        if (data.missing_headers && data.missing_headers.length > 0) {
            const list = document.createElement('ul');
            list.className = 'item-list';
            data.missing_headers.forEach(header => {
                list.innerHTML += `
                    <li class="list-item warning">
                        <i data-lucide="alert-circle"></i>
                        Missing: ${header}
                    </li>
                `;
            });
            headersContent.appendChild(list);
        } else {
            headersContent.innerHTML = `
                <div class="badge">A+ Rating</div>
                <ul class="item-list">
                    <li class="list-item good">
                        <i data-lucide="check-circle"></i>
                        All essential security headers present.
                    </li>
                </ul>
            `;
        }

        // 3. Sensitive Files
        filesContent.innerHTML = '';
        if (data.exposed_files && data.exposed_files.length > 0) {
            filesContent.innerHTML = `<div class="badge danger">CRITICAL FINDINGS</div>`;
            const list = document.createElement('ul');
            list.className = 'item-list';
            data.exposed_files.forEach(file => {
                list.innerHTML += `
                    <li class="list-item critical">
                        <i data-lucide="flame"></i>
                        Exposed File: ${file}
                    </li>
                `;
            });
            filesContent.appendChild(list);
        } else {
            filesContent.innerHTML = `
                <div class="badge">SECURE</div>
                <ul class="item-list">
                    <li class="list-item good">
                        <i data-lucide="shield-check"></i>
                        No common sensitive files exposed.
                    </li>
                </ul>
            `;
        }

        // Re-init icons for newly added HTML
        lucide.createIcons();

        // Show Results
        resultsGrid.classList.remove('hidden');
        // Small timeout to allow display:block to apply before animating opacity
        setTimeout(() => {
            resultsGrid.classList.add('visible');
            // Scroll to results smoothly if on small screen
            resultsGrid.scrollIntoView({ behavior: 'smooth', block: 'start' });
        }, 50);
    }
});
