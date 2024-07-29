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
