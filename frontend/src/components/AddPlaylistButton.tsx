import { useState } from "react";
import {
    Fab,
    Dialog,
    DialogTitle,
    DialogContent,
    TextField,
    DialogActions,
    Button,
} from "@mui/material";
import AddIcon from "@mui/icons-material/Add";
import axios from "axios";

interface AddPlaylistProps {
    user: string;
}

export default function AddPlaylist({ user }: AddPlaylistProps) {
    const [open, setOpen] = useState(false);
    const [link, setLink] = useState("");
    const [description, setDescription] = useState("");
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
        setLink("");
        setDescription("");
        setError("");
    };

    const handleAddPlaylist = async () => {
        if (!link.trim() || !description.trim()) {
            setError("Both link and description are required.");
            return;
        }

        try {
            setLoading(true);

            const token = sessionStorage.getItem("token");
            if (!token) {
                setError("You must be logged in to add a playlist.");
                setLoading(false);
                return;
            }


            const response = await axios.post(
                "http://localhost:8080/playlist/add", // Backend API endpoint
                {
                    username: user, // Pass the username received as a prop
                    link: link.trim(),  // Trimmed playlist link
                    description: description.trim(), // Trimmed description
                },
                {
                    headers: {
                        Authorization: `Bearer ${token}`, // Add Bearer token in the header
                    },
                }
            );

            console.log("Playlist added:", response.data);

            handleClose();
        } catch (err: any) {
            console.error("Error adding playlist:", err);
            setError("Failed to add playlist. Please try again.");
        } finally {
            setLoading(false);
        }
    };

    return (
        <div>
            { }
            <Fab
                color="primary"
                aria-label="add"
                onClick={handleClickOpen}
                style={{ position: "fixed", bottom: 16, right: 16 }}
            >
                <AddIcon />
            </Fab>

            {/* Dialog for Adding Playlist */}
            <Dialog open={open} onClose={handleClose}>
                <DialogTitle>Add New Playlist</DialogTitle>
                <DialogContent>
                    <TextField
                        autoFocus
                        margin="dense"
                        label="Link"
                        type="url"
                        fullWidth
                        variant="outlined"
                        value={link}
                        onChange={(e) => setLink(e.target.value)}
                        error={!!error && !link.trim()}
                        helperText={!!error && !link.trim() ? "Link is required" : ""}
                    />
                    <TextField
                        margin="dense"
                        label="Description"
                        type="text"
                        fullWidth
                        variant="outlined"
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                        error={!!error && !description.trim()}
                        helperText={
                            !!error && !description.trim()
                                ? "Description is required"
                                : ""
                        }
                    />
                    {error && <p style={{ color: "red" }}>{error}</p>}
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose} color="secondary" disabled={loading}>
                        Cancel
                    </Button>
                    <Button
                        onClick={handleAddPlaylist}
                        color="primary"
                        disabled={loading}
                    >
                        {loading ? "Adding..." : "Add"}
                    </Button>
                </DialogActions>
            </Dialog>
        </div>
    );
}
