import React from 'react';
import { Star } from 'lucide-react';
import MovieCard from '../components/MovieCard';

const RECOMMENDED = [
    { id: 7, title: "Tenet", rating: 7.3, releaseDate: "2020-08-22", poster: "https://image.tmdb.org/t/p/w500/k68nPLbIST6NP96JmTxmZijEvCA.jpg" },
    { id: 8, title: "Inception", rating: 8.8, releaseDate: "2010-07-15", poster: "https://image.tmdb.org/t/p/w500/oYuLEt3zVCKq57qu2F8dT7NIa6f.jpg" },
];

const Recommendations = () => {
    return (
        <div className="container animate-fade-in" style={{ padding: '32px 24px' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '16px', marginBottom: '32px' }}>
                <Star size={32} color="var(--warning)" fill="var(--warning)" />
                <h2 style={{ fontSize: '2.5rem' }}>For You</h2>
            </div>

            <p style={{ color: 'var(--text-muted)', marginBottom: '32px', fontSize: '1.1rem' }}>Based on your high ratings of sci-fi thrillers.</p>

            <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(220px, 1fr))', gap: '24px' }}>
                {RECOMMENDED.map(movie => (
                    <MovieCard key={movie.id} {...movie} />
                ))}
            </div>
        </div>
    );
};

export default Recommendations;
