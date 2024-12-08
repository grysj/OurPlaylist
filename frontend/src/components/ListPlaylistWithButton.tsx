import React, { useState } from "react";
import { Playlist } from "../types/types";
import LikePlaylistButton from "./LikePlaylistButton";

interface ListPlaylistsProps {
    playlists: Playlist[];
}

const ListPlaylists: React.FC<ListPlaylistsProps> = ({ playlists }) => {
    const [updatedPlaylists, setUpdatedPlaylists] = useState(playlists);

    const handleLikeUpdate = (playlistId: number) => {
        setUpdatedPlaylists((prev) =>
            prev.map((playlist) =>
                playlist.id === playlistId
                    ? { ...playlist, like_count: playlist.like_count + 1 }
                    : playlist
            )
        );
    };

    if (!Array.isArray(updatedPlaylists) || updatedPlaylists.length === 0) {
        return (
            <div>
                <p>No playlists</p>
            </div>
        );
    }

    return (
        <div>
            {updatedPlaylists.map((playlist) => (
                <div key={playlist.id}>
                    <a href={playlist.link}>{playlist.description}</a> LIKES: {playlist.like_count}
                    <LikePlaylistButton
                        playlistId={playlist.id}
                        onLike={() => handleLikeUpdate(playlist.id)}
                    />
                </div>
            ))}
        </div>
    );
};

export default ListPlaylists;
