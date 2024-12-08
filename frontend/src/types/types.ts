export interface Playlist {
    id: number;
    user_id: number;
    created_at: string;
    link: string;
    like_count: number;
    description: string;
}

export interface UserPlaylistsResponse {
    username: string;
    user_playlists: Playlist[];
    liked_playlists: Playlist[];
}