import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function Register() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [name, setName] = useState('');
    const navigate = useNavigate();

    const handleRegister = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email, password, name }),
            });

            const data = await response.json();

            if (response.ok) {
                // Registro exitoso, redirigir al usuario a la página de inicio de sesión
                navigate('/login');
            } else {
                // Mostrar mensaje de error al usuario
                console.error(data.error);
            }
        } catch (error) {
            console.error('Error al registrar:', error);
        }
    };

    return (
        <div className="container">
            <div className="form-container">
                <h1>Create an account</h1>
                <input
                    type="te"
                    className="input-field"
                    placeholder="Full Name"
                    value={name}
                    onChange={e => setName(e.target.value)}
                />
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
                <button className="button" onClick={handleRegister}>Register</button>
                <p className="link">Have an account already? <a href="/login">Sign in here</a></p>
            </div>
        </div>
    );
}

export default Register;
