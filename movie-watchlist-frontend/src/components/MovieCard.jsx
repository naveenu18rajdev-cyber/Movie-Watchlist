import React from 'react';
import { Star, PlusCircle } from 'lucide-react';

const MovieCard = ({ title, poster, rating, releaseDate }) => {
    return (
        <div className="glass-panel" style={{ overflow: 'hidden', display: 'flex', flexDirection: 'column', transition: 'transform 0.3s ease, box-shadow 0.3s ease', cursor: 'pointer' }}
            onMouseEnter={(e) => { e.currentTarget.style.transform = 'translateY(-8px)'; e.currentTarget.style.boxShadow = 'var(--shadow-glow)'; }}
            onMouseLeave={(e) => { e.currentTarget.style.transform = 'translateY(0)'; e.currentTarget.style.boxShadow = '0 4px 6px rgba(0,0,0,0.1)'; }}
        >
            <div style={{ position: 'relative', width: '100%', paddingTop: '150%', background: '#1e293b' }}>
                {poster ? (
                    <img src={poster} alt={title} style={{ position: 'absolute', top: 0, left: 0, width: '100%', height: '100%', objectFit: 'cover' }} />
                ) : (
                    <div style={{ position: 'absolute', top: 0, left: 0, width: '100%', height: '100%', display: 'flex', alignItems: 'center', justifyContent: 'center', color: 'var(--text-muted)' }}>No Image</div>
                )}

                <div style={{ position: 'absolute', top: '12px', right: '12px', background: 'rgba(0,0,0,0.7)', backdropFilter: 'blur(4px)', padding: '4px 8px', borderRadius: '12px', display: 'flex', alignItems: 'center', gap: '4px', fontSize: '0.85rem', fontWeight: 'bold' }}>
                    <Star size={14} color="var(--warning)" fill="var(--warning)" />
                    {rating || 'N/A'}
                </div>
            </div>

            <div style={{ padding: '16px', display: 'flex', flexDirection: 'column', flexGrow: 1 }}>
                <h3 style={{ fontSize: '1.1rem', marginBottom: '4px', display: '-webkit-box', WebkitLineClamp: 1, WebkitBoxOrient: 'vertical', overflow: 'hidden' }}>{title}</h3>
                <p style={{ color: 'var(--text-muted)', fontSize: '0.9rem', marginBottom: '16px' }}>{releaseDate?.split('-')[0] || 'Unknown Year'}</p>

                <button className="btn-outline" style={{ display: 'flex', alignItems: 'center', justifyContent: 'center', gap: '8px', padding: '8px', width: '100%', marginTop: 'auto' }}>
                    <PlusCircle size={18} /> Watchlist
                </button>
            </div>
        </div>
    );
};

export default MovieCard;
