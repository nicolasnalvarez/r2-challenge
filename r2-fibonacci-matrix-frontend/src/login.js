import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import './styles.css';

function Login() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    const handleLogin = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email, password }),
            });

            const data = await response.json();

            if (response.ok) {
                localStorage.setItem('accessToken', data.token)
                navigate('/matrix');
            } else {
                setErrorMessage(data.error);
            }
        } catch (error) {
            console.error('there was an error trying to login user:', error);
        }
    };

    return (
        <div className="container">
            <div className="form-container">
                <h1>Log In to start</h1>
                <input
                    type="email"
                    className="input-field"
                    placeholder="Email"
                    value={email}
                    onChange={e => setEmail(e.target.value)}
                />
                <input
                    type="password"
                    className="input-field"
                    placeholder="Password"
                    value={password}
                    onChange={e => setPassword(e.target.value)}
                />
                <button className="button" onClick={handleLogin}>Log In</button>
                {errorMessage && <p className="error-message">{errorMessage}</p>}
                <p className="link">Dont have an account? <a href="/register">Register here</a></p>
            </div>
        </div>
    );
}

export default Login;
