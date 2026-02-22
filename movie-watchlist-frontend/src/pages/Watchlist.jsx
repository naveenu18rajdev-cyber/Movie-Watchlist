import React, { useState } from 'react';
import { List, Filter } from 'lucide-react';
import MovieCard from '../components/MovieCard';

const MOCK_WATCHLIST = [
    { id: 2, title: "Oppenheimer", rating: 8.1, releaseDate: "2023-07-19", poster: "https://image.tmdb.org/t/p/w500/8Gxv8gSFCU0XGDykEGv7zR1n2ua.jpg", status: "Watched" },
    { id: 4, title: "Interstellar", rating: 8.4, releaseDate: "2014-11-05", poster: "https://image.tmdb.org/t/p/w500/gEU2QlsEOWepVNzMU5cR0yD2z2m.jpg", status: "Want to Watch" },
];

const Watchlist = () => {
    const [filter, setFilter] = useState('All');

    return (
        <div className="container animate-fade-in" style={{ padding: '32px 24px' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-end', marginBottom: '32px' }}>
                <div style={{ display: 'flex', alignItems: 'center', gap: '16px' }}>
                    <List size={32} color="var(--accent-primary)" />
                    <h2 style={{ fontSize: '2.5rem' }}>Your Watchlist</h2>
                </div>

                <div style={{ display: 'flex', gap: '12px' }}>
                    {['All', 'Want to Watch', 'Watching', 'Watched'].map(f => (
                        <button key={f} className={filter === f ? "btn-primary" : "btn-outline"} onClick={() => setFilter(f)} style={{ padding: '6px 16px', fontSize: '0.9rem' }}>
                            {f}
                        </button>
                    ))}
                </div>
            </div>

            <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(220px, 1fr))', gap: '24px' }}>
                {MOCK_WATCHLIST.filter(m => filter === 'All' || m.status === filter).map(movie => (
                    <div key={movie.id} style={{ position: 'relative' }}>
                        <MovieCard {...movie} />
                        <div style={{ position: 'absolute', top: '-10px', left: '-10px', background: movie.status === 'Watched' ? 'var(--success)' : 'var(--accent-secondary)', color: 'white', padding: '4px 12px', borderRadius: '12px', fontSize: '0.8rem', fontWeight: 'bold', zIndex: 10 }}>
                            {movie.status}
                        </div>
                    </div>
                ))}
            </div>

            {MOCK_WATCHLIST.length === 0 && (
                <div style={{ textAlign: 'center', padding: '64px', color: 'var(--text-muted)' }}>
                    <p>Your watchlist is currently empty.</p>
                </div>
            )}
        </div>
    );
};

export default Watchlist;
