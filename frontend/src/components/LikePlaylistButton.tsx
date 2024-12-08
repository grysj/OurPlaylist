import React, { useState } from "react";
import axios from "axios";

interface LikePlaylistProps {
    playlistId: number;
    onLike: () => void;
}

const LikePlaylistButton: React.FC<LikePlaylistProps> = ({ playlistId, onLike }) => {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | null>(null);

    const handleLike = async () => {
        try {
            setLoading(true);
            setError(null);

            const token = sessionStorage.getItem("token");
            if (!token) {
                setError("You must be logged in to like a playlist.");
                return;
            }

            await axios.post(
                "http://localhost:8080/playlist/like",
                {
                    playlist_id: playlistId,
                    username: sessionStorage.getItem("username")
                },
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                }
            );

            onLike();
        } catch (err) {
            console.error("Error liking playlist:", err);
            setError("Failed to like the playlist. Please try again.");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div>
            <button onClick={handleLike} disabled={loading}>
                {loading ? "Liking..." : "Like"}
            </button>
            {error && <p style={{ color: "red" }}>{error}</p>}
        </div>
    );
};

export default LikePlaylistButton;
