// Éléments DOM
const searchInput = document.getElementById('search-input');
const suggestionsBox = document.getElementById('suggestions-box');

// Debounce pour limiter les requêtes
let debounceTimer;

// Écouteur sur l'input
searchInput.addEventListener('input', function() {
    clearTimeout(debounceTimer);
    
    const query = this.value.trim();
    
    // Si moins de 2 caractères, cache les suggestions
    if (query.length < 2) {
        suggestionsBox.innerHTML = '';
        suggestionsBox.style.display = 'none';
        return;
    }
    
    // Délai de 300ms avant de lancer la recherche
    debounceTimer = setTimeout(() => {
        console.log('🔍 Recherche:', query);
        fetchSuggestions(query);
    }, 300);
});

// Fonction pour récupérer les suggestions
function fetchSuggestions(query) {
    fetch(`/search?q=${encodeURIComponent(query)}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Erreur serveur');
            }
            return response.json();
        })
        .then(data => {
            console.log('✅ Résultats:', data);
            displaySuggestions(data);
        })
        .catch(error => {
            console.error('❌ Erreur fetch:', error);
            suggestionsBox.innerHTML = '<div style="padding: 10px; color: #ff6b6b;">Erreur de recherche</div>';
        });
}

// Fonction pour afficher les suggestions
function displaySuggestions(suggestions) {
    // Vide la liste
    suggestionsBox.innerHTML = '';
    
    // Si aucun résultat
    if (!suggestions || suggestions.length === 0) {
        suggestionsBox.innerHTML = '<div style="padding: 10px; color: #888;">Aucun résultat</div>';
        suggestionsBox.style.display = 'block';
        return;
    }
    
    // Crée la liste des suggestions
    suggestions.forEach(item => {
        const div = document.createElement('div');
        div.className = 'suggestion-item';
        
        // Emoji selon le type
        const emoji = getEmoji(item.type);
        
        // Texte de la suggestion
        let text = `${emoji} ${item.value}`;
        if (item.type !== 'artist') {
            text += ` <span style="color: #888; font-size: 13px;">(${item.artist})</span>`;
        }
        
        div.innerHTML = text;
        
        // Clic sur une suggestion
        div.addEventListener('click', () => {
            console.log('✅ Sélection:', item);
            window.location.href = `/artist/${item.id}`;
        });
        
        suggestionsBox.appendChild(div);
    });
    
    suggestionsBox.style.display = 'block';
}

// Fonction pour obtenir l'emoji selon le type
function getEmoji(type) {
    const emojis = {
        'artist': '👨‍🎤',
        'member': '🎤',
        'location': '📍',
        'creation-date': '📅',
        'first-album': '💿'
    };
    return emojis[type] || '🔍';
}

// Fermer les suggestions si on clique ailleurs
document.addEventListener('click', function(e) {
    if (!searchInput.contains(e.target) && !suggestionsBox.contains(e.target)) {
        suggestionsBox.style.display = 'none';
    }
});

// Réafficher si on refocus l'input
searchInput.addEventListener('focus', function() {
    if (suggestionsBox.innerHTML && this.value.length >= 2) {
        suggestionsBox.style.display = 'block';
    }
});
