const searchInput = document.getElementById('search-input');
const suggestionsBox = document.getElementById('suggestions-box');

let debounceTimer;

searchInput.addEventListener('input', (e) => {
    clearTimeout(debounceTimer);
    const query = e.target.value.trim();
    
    if (query.length < 2) {
        suggestionsBox.style.display = 'none';
        return;
    }
    
    debounceTimer = setTimeout(() => {
        fetch(`/api/search?q=${encodeURIComponent(query)}`)
            .then(res => res.json())
            .then(suggestions => {
                displaySuggestions(suggestions);
            });
    }, 300);
});

function displaySuggestions(suggestions) {
    if (suggestions.length === 0) {
        suggestionsBox.style.display = 'none';
        return;
    }
    
    suggestionsBox.innerHTML = suggestions.map(s => `
        <div class="suggestion-item" onclick="selectSuggestion('${s.value}')">
            <div class="suggestion-type">${s.type}</div>
            <div>${s.value}</div>
        </div>
    `).join('');
    
    suggestionsBox.style.display = 'block';
}

function selectSuggestion(value) {
    searchInput.value = value;
    suggestionsBox.style.display = 'none';
    
    // Soumettre le formulaire de filtres
    document.getElementById('filters-form').submit();
}

// Cacher les suggestions si on clique ailleurs
document.addEventListener('click', (e) => {
    if (!searchInput.contains(e.target)) {
        suggestionsBox.style.display = 'none';
    }
});
