import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import Discovery from './pages/Discovery';
import Schedule from './pages/Schedule';
import Watchlist from './pages/Watchlist';
import Recommendations from './pages/Recommendations';
import News from './pages/News';

// Placeholder Pages
const Home = () => <div className="container animate-fade-in" style={{ marginTop: '48px', textAlign: 'center' }}>
  <h1 style={{ fontSize: '4rem', marginBottom: '16px' }} className="text-gradient">Discover Cinematic Brilliance</h1>
  <p style={{ fontSize: '1.25rem', color: 'var(--text-muted)', maxWidth: '600px', margin: '0 auto 32px' }}>
    Track, rate, and explore the universe of movies. Get tailored recommendations curated just for you.
  </p>
  <button className="btn-primary" style={{ fontSize: '1.1rem', padding: '16px 32px' }}>Explore Now</button>
</div>;

function App() {
  return (
    <Router>
      <div className="app-wrapper">
        <Navbar />
        <main>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/discovery" element={<Discovery />} />
            <Route path="/schedule" element={<Schedule />} />
            <Route path="/watchlist" element={<Watchlist />} />
            <Route path="/recommendations" element={<Recommendations />} />
            <Route path="/news" element={<News />} />
            <Route path="/auth" element={<div className="container"><h2>Auth (Coming Soon)</h2></div>} />
          </Routes>
        </main>
      </div>
    </Router>
  );
}

export default App;
