function createStars() {
    const starsContainer = document.getElementById('stars');
    const starsCount = 200;

    for (let i = 0; i < starsCount; i++) {
        const star = document.createElement('div');
        star.className = 'star';

        const size = Math.random() * 2;
        const posX = Math.random() * 100;
        const posY = Math.random() * 100;
        const duration = 2 + Math.random() * 3;
        const delay = Math.random() * 5;
        const opacity = 0.5 + Math.random() * 0.5;

        star.style.width = `${size}px`;
        star.style.height = `${size}px`;
        star.style.left = `${posX}%`;
        star.style.top = `${posY}%`;
        star.style.setProperty('--duration', `${duration}s`);
        star.style.setProperty('--opacity', opacity);
        star.style.animationDelay = `${delay}s`;

        starsContainer.appendChild(star);
    }

    setInterval(() => {
        const meteor = document.createElement('div');
        meteor.className = 'meteor';

        const startX = Math.random() * 100;
        const startY = Math.random() * 30;
        const duration = 0.5 + Math.random() * 1;

        meteor.style.left = `${startX}%`;
        meteor.style.top = `${startY}%`;
        meteor.style.animationDuration = `${duration}s`;

        document.body.appendChild(meteor);

        setTimeout(() => {
            meteor.remove();
        }, duration * 1000);
    }, 2000);
}

document.addEventListener('DOMContentLoaded', createStars);