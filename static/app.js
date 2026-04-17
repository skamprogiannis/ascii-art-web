(() => {
    const MAX_CHARS = 1000;
    const WARN_THRESHOLD = 800;

    const textarea = document.getElementById('text');
    const counter = document.getElementById('char-counter');
    const form = document.getElementById('art-form');
    const warning = document.getElementById('empty-warning');
    const outputPanel = document.querySelector('.panel-output');

    if (!textarea || !counter || !form || !warning || !outputPanel) {
        return;
    }

    const picker = document.getElementById('color-picker');
    const colorTx = document.getElementById('color-text');
    const swatch = document.getElementById('color-swatch');
    const substrInput = document.getElementById('substr');

    let fetchAbortController = null;
    let typingTimer;
    let substrTimer;

    const existingArtOutput = document.getElementById('art-output');
    let preferLightCanvas = !!(existingArtOutput && existingArtOutput.classList.contains('canvas-light'));

    const emptyStateHTML = `
        <div class="output-empty">
            <div class="output-empty-icon" aria-hidden="true">⌨</div>
            <p>Your ASCII art will appear here</p>
        </div>`;

    const escapeHTML = (value) => value
        .replaceAll('&', '&amp;')
        .replaceAll('<', '&lt;')
        .replaceAll('>', '&gt;')
        .replaceAll('"', '&quot;')
        .replaceAll("'", '&#39;');

    const parseDownloadFilename = (contentDisposition, fallbackName) => {
        if (!contentDisposition) {
            return fallbackName;
        }

        const utf8Match = contentDisposition.match(/filename\*\s*=\s*UTF-8''([^;]+)/i);
        if (utf8Match && utf8Match[1]) {
            return decodeURIComponent(utf8Match[1]);
        }

        const plainMatch = contentDisposition.match(/filename\s*=\s*"?(?<name>[^";]+)"?/i);
        if (plainMatch && plainMatch.groups && plainMatch.groups.name) {
            return plainMatch.groups.name;
        }

        return fallbackName;
    };

    const flashButtonState = (button, text, className = '') => {
        if (!button) {
            return;
        }

        const originalText = button.dataset.defaultLabel || button.textContent;
        button.dataset.defaultLabel = originalText;
        button.textContent = text;
        if (className) {
            button.classList.add(className);
        }
        setTimeout(() => {
            button.textContent = originalText;
            if (className) {
                button.classList.remove(className);
            }
        }, 2000);
    };

    const updateCounter = () => {
        const n = textarea.value.length;
        counter.textContent = `${n} / ${MAX_CHARS}`;
        counter.classList.toggle('warn', n > WARN_THRESHOLD);
        counter.classList.toggle('over', n >= MAX_CHARS);
    };

    const updateSwatch = (value) => {
        if (swatch) {
            swatch.style.background = value || 'transparent';
        }
    };

    const hasRenderableOutput = (payload) => {
        if (!payload) {
            return false;
        }
        return Boolean(payload.error || payload.ascii_art || payload.ascii_art_html);
    };

    const renderOutput = (payload) => {
        if (!hasRenderableOutput(payload)) {
            outputPanel.innerHTML = emptyStateHTML;
            return;
        }

        const hasArt = Boolean(payload.ascii_art || payload.ascii_art_html);
        const canvasClass = preferLightCanvas ? 'canvas-light' : 'canvas-dark';
        const artHTML = payload.ascii_art_html || escapeHTML(payload.ascii_art || '');
        const errorHTML = payload.error
            ? `<div class="error-box" role="alert">${escapeHTML(payload.error)}</div>`
            : '';

        outputPanel.innerHTML = `
            <div class="result-header">
                <h2>Output</h2>
                <div class="output-actions">
                    <button type="button" id="theme-toggle" class="copy-btn" title="Toggle Dark/Light Canvas">◑ Background</button>
                    ${hasArt
                        ? `<span class="badge badge-success">✓ Generated</span>
                           <button class="copy-btn" id="copy-btn" type="button">⎘ Copy</button>
                           <select id="export-format" class="copy-btn export-format" aria-label="Export format">
                               <option value="txt">TXT</option>
                               <option value="html">HTML</option>
                               <option value="json">JSON</option>
                           </select>
                           <button class="copy-btn" id="download-btn" type="button" title="Download exported file">📥 Export File</button>`
                        : `<span class="badge badge-error">✕ Error</span>`}
                </div>
            </div>
            ${errorHTML}
            ${hasArt
                ? `<div class="pre-wrap"><pre id="art-output" class="${canvasClass}">${artHTML}</pre></div>`
                : ''}`;

        bindOutputEvents();
    };

    const getCurrentArtText = () => {
        const artOutput = document.getElementById('art-output');
        return artOutput ? artOutput.textContent : '';
    };

    function bindOutputEvents() {
        const copyBtn = document.getElementById('copy-btn');
        if (copyBtn) {
            copyBtn.addEventListener('click', () => {
                navigator.clipboard.writeText(getCurrentArtText()).then(() => {
                    copyBtn.textContent = '✓ Copied!';
                    copyBtn.classList.add('copied');
                    setTimeout(() => {
                        copyBtn.textContent = '⎘ Copy';
                        copyBtn.classList.remove('copied');
                    }, 2000);
                });
            });
        }

        const downloadBtn = document.getElementById('download-btn');
        if (downloadBtn) {
            downloadBtn.addEventListener('click', async () => {
                const exportFormat = document.getElementById('export-format');
                const format = exportFormat ? exportFormat.value : 'txt';
                const defaultFilename = `ascii-art.${format}`;

                downloadBtn.disabled = true;

                try {
                    const formData = new FormData(form);
                    formData.set('format', format);

                    const response = await fetch('/export', {
                        method: 'POST',
                        body: new URLSearchParams(formData),
                    });

                    if (!response.ok) {
                        const message = (await response.text()) || `${response.status} ${response.statusText}`;
                        throw new Error(message);
                    }

                    const blob = await response.blob();
                    const filename = parseDownloadFilename(response.headers.get('Content-Disposition'), defaultFilename);
                    const url = URL.createObjectURL(blob);
                    const link = document.createElement('a');
                    link.href = url;
                    link.download = filename;
                    link.click();
                    URL.revokeObjectURL(url);
                    flashButtonState(downloadBtn, '✓ Exported!', 'copied');
                } catch (err) {
                    console.error('Failed to export art:', err);
                    flashButtonState(downloadBtn, '✕ Export failed');
                } finally {
                    downloadBtn.disabled = false;
                }
            });
        }

        const themeToggle = document.getElementById('theme-toggle');
        const artOutput = document.getElementById('art-output');
        if (themeToggle && artOutput) {
            themeToggle.addEventListener('click', () => {
                const isDark = artOutput.classList.contains('canvas-dark');
                artOutput.classList.toggle('canvas-dark', !isDark);
                artOutput.classList.toggle('canvas-light', isDark);
                preferLightCanvas = artOutput.classList.contains('canvas-light');
            });
        }
    }

    const submitIfTextPresent = () => {
        if (textarea.value.trim()) {
            form.requestSubmit();
        }
    };

    document.addEventListener('keydown', (event) => {
        if (event.key === 'Enter' && document.activeElement !== textarea) {
            event.preventDefault();
            form.requestSubmit();
        }
    });

    document.querySelectorAll('.style-card input[type=radio]').forEach((radio) => {
        radio.addEventListener('change', () => {
            document.querySelectorAll('.style-card').forEach((card) => card.classList.remove('active'));
            radio.closest('.style-card').classList.add('active');
            submitIfTextPresent();
        });
    });

    if (picker && colorTx) {
        picker.addEventListener('input', () => {
            colorTx.value = picker.value;
            updateSwatch(picker.value);
        });

        colorTx.addEventListener('input', () => {
            const value = colorTx.value.trim();
            updateSwatch(value);
            if (/^#[0-9a-fA-F]{6}$/.test(value)) {
                picker.value = value;
            }
        });

        picker.addEventListener('change', submitIfTextPresent);
        colorTx.addEventListener('change', submitIfTextPresent);
        updateSwatch(colorTx.value || '#8b5cf6');
    }

    document.querySelectorAll('#color-presets button').forEach((button) => {
        button.addEventListener('click', () => {
            const color = button.dataset.color;
            if (colorTx) {
                colorTx.value = color;
            }
            if (picker) {
                picker.value = color;
            }
            updateSwatch(color);
            document.querySelectorAll('#color-presets button').forEach((b) => b.classList.remove('active'));
            button.classList.add('active');
            submitIfTextPresent();
        });
    });

    form.addEventListener('submit', async (event) => {
        if (!textarea.value.trim()) {
            event.preventDefault();
            warning.classList.add('visible');
            textarea.focus();
            if (fetchAbortController) {
                fetchAbortController.abort();
            }
            outputPanel.innerHTML = emptyStateHTML;
            return;
        }

        event.preventDefault();
        warning.classList.remove('visible');
        outputPanel.style.opacity = '0.5';

        if (fetchAbortController) {
            fetchAbortController.abort();
        }
        const currentController = new AbortController();
        fetchAbortController = currentController;

        try {
            const formData = new FormData(form);
            const response = await fetch(form.action, {
                method: form.method,
                headers: {
                    'Accept': 'application/json',
                },
                body: new URLSearchParams(formData),
                signal: currentController.signal,
            });

            const contentType = (response.headers.get('Content-Type') || '').toLowerCase();
            let payload;
            if (contentType.includes('application/json')) {
                payload = await response.json();
            } else {
                const text = await response.text();
                payload = { error: text || 'Unexpected server response.' };
            }

            if (!response.ok && !payload.error) {
                payload.error = `${response.status} ${response.statusText}`;
            }

            renderOutput(payload);
        } catch (err) {
            if (err.name === 'AbortError') {
                return;
            }
            console.error('Failed to generate art:', err);
            renderOutput({ error: 'Request failed. Please try again.' });
        } finally {
            if (!currentController.signal.aborted) {
                outputPanel.style.opacity = '1';
            }
        }
    });

    textarea.addEventListener('input', () => {
        warning.classList.remove('visible');
        updateCounter();
        clearTimeout(typingTimer);
        if (textarea.value.trim()) {
            typingTimer = setTimeout(() => form.requestSubmit(), 500);
            return;
        }
        if (fetchAbortController) {
            fetchAbortController.abort();
        }
        outputPanel.innerHTML = emptyStateHTML;
    });

    if (substrInput) {
        substrInput.addEventListener('input', () => {
            clearTimeout(substrTimer);
            if (textarea.value.trim()) {
                substrTimer = setTimeout(() => form.requestSubmit(), 500);
            }
        });
    }

    bindOutputEvents();
    updateCounter();

    const canvas = document.getElementById('matrix-bg');
    if (!canvas) {
        return;
    }
    const ctx = canvas.getContext('2d');
    if (!ctx) {
        return;
    }

    let width;
    let height;
    let columns;
    let drops;
    const fontSize = 14;
    const letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789$+-*/=%""\'#&_(),.;:?!\\|{}<>[]^~';

    const initMatrix = () => {
        width = canvas.width = window.innerWidth;
        height = canvas.height = window.innerHeight;
        columns = Math.floor(width / fontSize);
        drops = [];
        for (let x = 0; x < columns; x++) {
            drops[x] = 1;
        }
    };

    const drawMatrix = () => {
        ctx.fillStyle = 'rgba(8, 11, 20, 0.1)';
        ctx.fillRect(0, 0, width, height);
        ctx.fillStyle = '#0f0';
        ctx.font = `${fontSize}px monospace`;

        for (let i = 0; i < drops.length; i++) {
            const text = letters.charAt(Math.floor(Math.random() * letters.length));
            ctx.fillText(text, i * fontSize, drops[i] * fontSize);
            if (drops[i] * fontSize > height && Math.random() > 0.975) {
                drops[i] = 0;
            }
            drops[i]++;
        }
    };

    initMatrix();
    setInterval(drawMatrix, 50);
    window.addEventListener('resize', initMatrix);
})();
