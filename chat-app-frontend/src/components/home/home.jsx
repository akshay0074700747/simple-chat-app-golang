// HomePage.js

import React from 'react';
import './home.css';

const HomePage = () => {
  return (
    <div className="home-page">
      <header className="header">
        <h1>Welcome to My Beautiful Website</h1>
        <p>Your go-to destination for amazing content</p>
      </header>
      <section className="content">
        <div className="feature">
          <h2>Discover</h2>
          <p>Explore a world of exciting opportunities and experiences.</p>
        </div>
        <div className="feature">
          <h2>Create</h2>
          <p>Unleash your creativity and bring your ideas to life.</p>
        </div>
        <div className="feature">
          <h2>Connect</h2>
          <p>Join our vibrant community and stay connected.</p>
        </div>
      </section>
    </div>
  );
};

export default HomePage;
