document.addEventListener("DOMContentLoaded", () => {
  const form = document.getElementById("filters-form");
  const resetBtn = document.getElementById("reset-filters");

  if (!form || !resetBtn) return;

  resetBtn.addEventListener("click", () => {
    // Reset tous les inputs
    form.querySelectorAll("input").forEach((input) => {
      if (input.type === "checkbox") {
        input.checked = false;
      } else {
        input.value = "";
      }
    });

    // Optionnel : recharger la page sans paramètres
    window.location.href = "/";
  });
});
