import { useEffect, useState } from "react";
import Menu from "../components/Menu";
import axios from "axios";
import { Playlist, UserPlaylistsResponse } from "../types/types";
import ListPlaylists from "../components/ListPlaylist";
import LoggedMenu from "../components/LoggedMenu";
import AddPlaylistButton from "../components/AddPlaylistButton";


export default function User() {
    const [username, setUsername] = useState("");
    const [userPlaylists, setUserPlaylists] = useState<Playlist[]>([]);
    const [likedPlaylists, setLikedPlaylists] = useState<Playlist[]>([]);
    const [error, setError] = useState("");

    useEffect(() => {
        const fetchData = async () => {
            try {
                const token = sessionStorage.getItem("token");

                const response = await axios.get<UserPlaylistsResponse>("http://localhost:8080/playlists", {
                    headers: {
                        Authorization: `Bearer ${token}`,
                    },
                });
                setUsername(response.data.username);
                setUserPlaylists(response.data.user_playlists || []);
                setLikedPlaylists(response.data.liked_playlists || []);
            } catch (err: any) {
                console.error("Error fetching playlists:", err);
                setError("Failed to fetch playlists. Please try again.");
            }
        };

        fetchData();
    }, []);

    return (
        <>
            <LoggedMenu />
            <div>
                <h1>Hello {username} !</h1>
                {error ? (
                    <p>{error}</p>
                ) : (
                    <>
                        <h2>Your Playlists</h2>
                        <ListPlaylists playlists={userPlaylists} />
                        <h2>Liked Playlists</h2>
                        <ListPlaylists playlists={likedPlaylists} />
                    </>
                )}
            </div>
            <AddPlaylistButton user={username} />
        </>
    );
}
