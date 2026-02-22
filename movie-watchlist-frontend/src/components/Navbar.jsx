import React from 'react';
import { Link } from 'react-router-dom';
import { Film, User, Search, List, Calendar, Star } from 'lucide-react';

const Navbar = () => {
    return (
        <nav className="glass-panel" style={{ margin: '16px auto', maxWidth: '1200px', display: 'flex', justifyContent: 'space-between', alignItems: 'center', padding: '16px 32px' }}>
            <Link to="/" style={{ display: 'flex', alignItems: 'center', gap: '12px' }}>
                <div style={{ background: 'var(--accent-gradient)', padding: '8px', borderRadius: '12px', display: 'flex' }}>
                    <Film size={24} color="#fff" />
                </div>
                <h1 style={{ fontSize: '1.5rem', letterSpacing: '-0.5px' }} className="text-gradient">CineVault</h1>
            </Link>

            <div style={{ display: 'flex', gap: '24px', alignItems: 'center', fontWeight: '500' }}>
                <Link to="/discovery" style={{ display: 'flex', alignItems: 'center', gap: '8px', opacity: 0.8 }}><Search size={18} /> Discover</Link>
                <Link to="/schedule" style={{ display: 'flex', alignItems: 'center', gap: '8px', opacity: 0.8 }}><Calendar size={18} /> Schedule</Link>
                <Link to="/watchlist" style={{ display: 'flex', alignItems: 'center', gap: '8px', opacity: 0.8 }}><List size={18} /> Watchlist</Link>
                <Link to="/recommendations" style={{ display: 'flex', alignItems: 'center', gap: '8px', opacity: 0.8 }}><Star size={18} /> For You</Link>
            </div>

            <div style={{ display: 'flex', gap: '16px' }}>
                <Link to="/auth" className="btn-outline">Log In</Link>
                <Link to="/auth?mode=register" className="btn-primary">Sign Up</Link>
            </div>
        </nav>
    );
};

export default Navbar;
