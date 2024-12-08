import React from "react";
import { Playlist } from "../types/types";



interface ListPlaylistsProps {
    playlists: Playlist[];
}

const ListPlaylists: React.FC<ListPlaylistsProps> = ({ playlists }) => {
    if (!Array.isArray(playlists) || playlists.length === 0) {
        return (
            <div>
                <p>No playlists</p>
            </div>
        );
    }

    return (
        <div>
            {playlists.map((playlist) => (
                <div key={playlist.id}>
                    <a href={playlist.link}>{playlist.description}</a> LIKES: {playlist.like_count}
                </div>
            ))}
        </div>
    );
};

export default ListPlaylists;
