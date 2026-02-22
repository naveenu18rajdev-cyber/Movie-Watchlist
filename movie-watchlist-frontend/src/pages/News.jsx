import React from 'react';
import { Newspaper } from 'lucide-react';

const NEWS_ARTICLES = [
    { id: 1, title: "Christopher Nolan announces new project targeting 2026 release", date: "2 Hours Ago", summary: "Following the massive success of Oppenheimer, the visionary director has revealed hints about his next film." },
    { id: 2, title: "Dune: Part Three officially greenlit by Legendary", date: "5 Hours Ago", summary: "Denis Villeneuve will return to direct the adaptation of Dune: Messiah, concluding Paul Atreides' arc." },
];

const News = () => {
    return (
        <div className="container animate-fade-in" style={{ padding: '32px 24px', maxWidth: '800px' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '16px', marginBottom: '32px' }}>
                <Newspaper size={32} color="var(--accent-secondary)" />
                <h2 style={{ fontSize: '2.5rem' }}>Entertainment News</h2>
            </div>

            <div style={{ display: 'flex', flexDirection: 'column', gap: '24px' }}>
                {NEWS_ARTICLES.map(article => (
                    <div key={article.id} className="glass-panel" style={{ padding: '24px', transition: 'transform 0.2s', cursor: 'pointer' }}
                        onMouseEnter={(e) => e.currentTarget.style.transform = 'scale(1.02)'}
                        onMouseLeave={(e) => e.currentTarget.style.transform = 'scale(1)'}>
                        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', marginBottom: '12px' }}>
                            <h3 style={{ fontSize: '1.25rem', color: 'var(--accent-primary)' }}>{article.title}</h3>
                            <span style={{ fontSize: '0.85rem', color: 'var(--text-muted)', whiteSpace: 'nowrap', marginLeft: '16px' }}>{article.date}</span>
                        </div>
                        <p style={{ color: 'var(--text-main)', opacity: 0.9 }}>{article.summary}</p>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default News;
