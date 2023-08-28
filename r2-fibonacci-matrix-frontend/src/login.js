import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';

function Login() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const history = useHistory();

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
                // Inicio de sesión exitoso, redirigir al usuario a la página principal
                history.push('/');
            } else {
                // Mostrar mensaje de error al usuario
                console.error(data.error);
            }
        } catch (error) {
            console.error('Error al iniciar sesión:', error);
        }
    };

    return (
        <div>
            <h1>Iniciar Sesión</h1>
            <input
                type="email"
                placeholder="Correo Electrónico"
                value={email}
                onChange={e => setEmail(e.target.value)}
            />
            <input
                type="password"
                placeholder="Contraseña"
                value={password}
                onChange={e => setPassword(e.target.value)}
            />
            <button onClick={handleLogin}>Iniciar Sesión</button>
            <p>No tienes cuenta? <a href="/register">Regístrate aquí</a></p>
        </div>
    );
}

export default Login;
