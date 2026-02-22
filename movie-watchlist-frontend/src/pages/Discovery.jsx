import React, { useState } from 'react';
import { Search } from 'lucide-react';
import MovieCard from '../components/MovieCard';

// Dummy data for visual presentation since backend won't have pre-loaded data locally yet
const SAMPLE_MOVIES = [
    { id: 1, title: "Dune: Part Two", rating: 8.8, releaseDate: "2024-03-01", poster: "https://image.tmdb.org/t/p/w500/1pdfLvkbY9ohJlCjQH2JGjjc91p.jpg" },
    { id: 2, title: "Oppenheimer", rating: 8.1, releaseDate: "2023-07-19", poster: "https://image.tmdb.org/t/p/w500/8Gxv8gSFCU0XGDykEGv7zR1n2ua.jpg" },
    { id: 3, title: "The Dark Knight", rating: 8.5, releaseDate: "2008-07-16", poster: "https://image.tmdb.org/t/p/w500/qJ2tW6WMUDux911r6m7haRef0WH.jpg" },
    { id: 4, title: "Interstellar", rating: 8.4, releaseDate: "2014-11-05", poster: "https://image.tmdb.org/t/p/w500/gEU2QlsEOWepVNzMU5cR0yD2z2m.jpg" },
    { id: 5, title: "Avatar: The Way of Water", rating: 7.6, releaseDate: "2022-12-14", poster: "https://image.tmdb.org/t/p/w500/t6HIqrHezINNdIExjkB7bEdW3c8.jpg" },
];

const Discovery = () => {
    const [query, setQuery] = useState('');

    return (
        <div className="container animate-fade-in" style={{ padding: '32px 24px' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-end', marginBottom: '32px' }}>
                <div>
                    <h2 style={{ fontSize: '2.5rem', marginBottom: '8px' }}>Discover Movies</h2>
                    <p style={{ color: 'var(--text-muted)' }}>Find your next favorite film</p>
                </div>

                <div className="glass-panel" style={{ display: 'flex', alignItems: 'center', padding: '8px 16px', borderRadius: '32px', width: '350px' }}>
                    <Search size={20} color="var(--text-muted)" />
                    <input
                        type="text"
                        placeholder="Search by title..."
                        value={query}
                        onChange={(e) => setQuery(e.target.value)}
                        style={{ background: 'transparent', border: 'none', outline: 'none', color: '#fff', marginLeft: '12px', width: '100%', fontSize: '1rem' }}
                    />
                </div>
            </div>

            <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(220px, 1fr))', gap: '24px' }}>
                {SAMPLE_MOVIES.map(movie => (
                    <MovieCard key={movie.id} {...movie} />
                ))}
            </div>
        </div>
    );
};

export default Discovery;
