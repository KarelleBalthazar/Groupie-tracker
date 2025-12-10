// Neon click effect
document.addEventListener("click", (e) => {
    if (e.target.classList.contains("neon-btn")) {
        e.target.classList.add("pulse");
        setTimeout(() => e.target.classList.remove("pulse"), 250);
    }
});
