import React, { useEffect } from 'react';
import { useHistory, useNavigate } from 'react-router-dom';

const withAuthorization = (WrappedComponent) => {
    return function WithAuthorization(props) {
        const navigate = useNavigate();

        useEffect(() => {
            // Verificar si el usuario está autenticado aquí
            // TODO seguir aca
            const isAuthenticated = /* Lógica para verificar la autenticación */;

            if (!isAuthenticated) {
                navigate('/login');
            }
        }, []);

        return <WrappedComponent {...props} />;
    };
};

export default withAuthorization;
