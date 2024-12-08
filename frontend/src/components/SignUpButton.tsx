import * as React from 'react';
import { useState } from 'react';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import axios from 'axios';

export const SignUpButton = () => {
    const [open, setOpen] = useState(false);
    const [formData, setFormData] = useState({
        mail: '',
        username: '',
        password: '',
    });
    const [response, setResponse] = useState("")

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData((prevData) => ({
            ...prevData,
            [name]: value
        }));
    };

    const handleSubmit = async () => {
        console.log('Form Data:', formData);
        try {
            const result = await axios.post('http://localhost:8080/users', formData);
            setResponse("An account was successfully created!");
            console.log("API Response:", result.data);
        } catch (error: any) {
            console.error("API Error:", error);
            setResponse("Something went wrong: " + (error.response?.data?.message || error.message));
        }

    };

    return (
        <React.Fragment>
            <Button
                variant="outlined"
                onClick={handleClickOpen}
                sx={{
                    color: 'black',
                    backgroundColor: 'white'
                }}
            >
                Sign up
            </Button>
            <Dialog open={open} onClose={handleClose}>
                <DialogTitle>Sign up</DialogTitle>
                <DialogContent>
                    <TextField
                        autoFocus
                        required
                        margin="dense"
                        id="email"
                        name="mail"
                        label="Email Address"
                        type="email"
                        fullWidth
                        variant="standard"
                        value={formData.mail}
                        onChange={handleChange}
                    />
                    <TextField
                        required
                        margin="dense"
                        id="username"
                        name="username"
                        label="Username"
                        type="text"
                        fullWidth
                        variant="standard"
                        value={formData.username}
                        onChange={handleChange}
                    />
                    <TextField
                        required
                        margin="dense"
                        id="password"
                        name="password"
                        label="Password"
                        type="password"
                        fullWidth
                        variant="standard"
                        value={formData.password}
                        onChange={handleChange}
                    />
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}>Cancel</Button>
                    <Button onClick={handleSubmit}>Sign up</Button>
                </DialogActions>
                <div>
                    {response}
                </div>
            </Dialog>
        </React.Fragment>
    );
};