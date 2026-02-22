import React from 'react';
import { Calendar, Clock, Film } from 'lucide-react';
import MovieCard from '../components/MovieCard';

const SCHEDULED_MOVIES = [
    { id: 10, title: "Deadpool 3", releaseDate: "2024-07-26", poster: "https://image.tmdb.org/t/p/w500/8cdWjvZQUExUUTzyp4t6EDMubfO.jpg" },
    { id: 11, title: "Gladiator 2", releaseDate: "2024-11-22", poster: "https://image.tmdb.org/t/p/w500/1p5xQv7dZ4i1lZ0D0v3f3eS2FQU.jpg" },
];

const Schedule = () => {
    return (
        <div className="container animate-fade-in" style={{ padding: '32px 24px' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '16px', marginBottom: '32px' }}>
                <Calendar size={32} color="var(--accent-primary)" />
                <h2 style={{ fontSize: '2.5rem' }}>Release Schedule</h2>
            </div>

            <div style={{ borderLeft: '2px solid var(--glass-border)', paddingLeft: '32px', position: 'relative' }}>
                <div style={{ position: 'absolute', left: '-11px', top: '0', background: 'var(--bg-primary)', padding: '4px' }}>
                    <Clock size={20} color="var(--accent-secondary)" />
                </div>

                <h3 style={{ fontSize: '1.5rem', marginBottom: '16px', color: 'var(--accent-primary)' }}>Coming Soon (2024)</h3>

                <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(220px, 1fr))', gap: '24px', marginBottom: '48px' }}>
                    {SCHEDULED_MOVIES.map(movie => (
                        <MovieCard key={movie.id} {...movie} />
                    ))}
                </div>
            </div>
        </div>
    );
};

export default Schedule;
