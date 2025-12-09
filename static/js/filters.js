document.addEventListener("DOMContentLoaded", () => {
  const form = document.getElementById("filters-form");
  const resetBtn = document.getElementById("reset-filters");

  if (!form || !resetBtn) return;

  resetBtn.addEventListener("click", () => {
    form.querySelectorAll("input").forEach((input) => {
      if (input.type === "checkbox") {
        input.checked = false;
      } else {
        input.value = "";
      }
    });

    // Recharger la page sans paramètres
    window.location.href = "/";
  });
});
