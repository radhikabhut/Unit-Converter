document.addEventListener('DOMContentLoaded', function () {
    const fromUnitSelect = document.getElementById('fromUnit');
    const toUnitSelect = document.getElementById('toUnit');
    const valueInput = document.getElementById('value'); // Ensure input has an ID

    // Restore stored values
    if (localStorage.getItem("fromUnit")) {
        fromUnitSelect.value = localStorage.getItem("fromUnit");
    }
    if (localStorage.getItem("toUnit")) {
        toUnitSelect.value = localStorage.getItem("toUnit");
    }
    if (localStorage.getItem("value")) {
        valueInput.value = localStorage.getItem("value");
    }

    fromUnitSelect.addEventListener('change', function () {
        const selectedCategory = fromUnitSelect.options[fromUnitSelect.selectedIndex].dataset.category;
        localStorage.setItem("fromUnit", fromUnitSelect.value);
        
        // Clear To Unit dropdown
        toUnitSelect.innerHTML = '<option value="">Select To Unit</option>';

        // Populate To Unit dropdown with compatible units
        Array.from(fromUnitSelect.options).forEach(option => {
            if (option.dataset.category === selectedCategory && option.value !== fromUnitSelect.value) {
                const newOption = document.createElement('option');
                newOption.value = option.value;
                newOption.textContent = option.textContent;

                if (option.value === localStorage.getItem("toUnit")) {
                    newOption.selected = true;
                }

                toUnitSelect.appendChild(newOption);
            }
        });
    });

    toUnitSelect.addEventListener('change', function () {
        localStorage.setItem("toUnit", toUnitSelect.value);
    });

    valueInput.addEventListener('input', function () {
        localStorage.setItem("value", valueInput.value);
    });
});
