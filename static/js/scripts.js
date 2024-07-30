document.addEventListener('DOMContentLoaded', function() {
    const features = document.querySelectorAll('.feature');
    let currentFeatureIndex = 0;

    function showFeature(index) {
        features.forEach((feature, i) => {
            feature.classList.toggle('active', i === index);
        });
    }

    function nextFeature() {
        currentFeatureIndex = (currentFeatureIndex + 1) % features.length;
        showFeature(currentFeatureIndex);
    }

    showFeature(currentFeatureIndex);
    setInterval(nextFeature, 5000);
});
document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('loginForm').addEventListener('submit', async function(e) {
        e.preventDefault(); // Prevent the default form submission

        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;
        const authKey = document.getElementById('auth-key').value;
        const userType = document.querySelector('input[name="user-type"]:checked').value;

        const response = await fetch('/admin', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, password, authKey, userType })
        });

        const data = await response.json(); // Parse the response as JSON
        if (response.ok) {
            if (data.userType === 'patient') {
                window.location.href = '/';
            } else if (data.userType === 'facility') {
                window.location.href = '/facility-dashboard';
            }
        } else {
            document.getElementById('message').textContent = data.message; // Display error message
        }
    });
});
