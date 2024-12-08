import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import { Playlist } from "../types/types";
import LoggedMenu from "../components/LoggedMenu";
import ListPlaylistsWithButton from "../components/ListPlaylistWithButton";


interface ProfileResponse {
    username: string;
    user_playlists: Playlist[];
}

export default function Profile() {
    const { id } = useParams<{ id: string }>();
    const [username, setUsername] = useState("");
    const [userPlaylists, setUserPlaylists] = useState<Playlist[]>([]);
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const token = sessionStorage.getItem("token");

                const response = await axios.get<ProfileResponse>(
                    `http://localhost:8080/profile`,
                    {
                        params: { username: id },
                        headers: {
                            Authorization: `Bearer ${token}`,
                        },
                    }
                );



                setUsername(response.data.username);
                setUserPlaylists(response.data.user_playlists || []);
                setLoading(false);
            } catch (err: any) {
                console.error("Error fetching profile data:", err);
                setError("Failed to fetch profile data. Please try again.");
                setLoading(false);
            }
        };

        if (id) fetchData();
    }, [id]);

    return (

        <div>
            <LoggedMenu />
            {loading ? (
                <p>Loading...</p>
            ) : error ? (
                <p>{error}</p>
            ) : (
                <>
                    <h1>{username}</h1>
                    <h2>{username} playlists</h2>
                    <ListPlaylistsWithButton playlists={userPlaylists} />
                </>
            )}
        </div>
    );
}
