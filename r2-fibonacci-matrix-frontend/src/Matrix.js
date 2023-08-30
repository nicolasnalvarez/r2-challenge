import React, { useState } from 'react';
import './matrix.css'

function Matrix() {
    const [rows, setRows] = useState('');
    const [cols, setCols] = useState('');
    const [matrix, setMatrix] = useState([]);
    const [loading, setLoading] = useState(false);
    const [errorMessage, setErrorMessage] = useState('');

    const handleCalculate = () => {
        setLoading(true);
        setErrorMessage('');
        fetch(`http://localhost:8080/api/matrix?rows=${rows}&columns=${cols}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + localStorage.getItem('accessToken'),
            },
        })
            .then(response => response.json()) // check if 401 to refresh token
            .then(data => {
                if (data.error) {
                    setErrorMessage(data.error);
                    setMatrix([]);
                } else {
                    setMatrix(data.rows);
                }
            })
            .finally(() => setLoading(false));
    };

    return (
        <div className="App">
            <div className="header">
                <h1 className="title">Fibonacci Spiral</h1>
                <p className="subtitle">Matrix properties</p>
                <div className="input-group">
                    <div className="gray-cell">Number of Rows</div>
                    <input
                        type="number"
                        value={rows}
                        onChange={e => setRows(e.target.value)}
                        className="editable-input"
                    />
                    <div className="gray-cell">Number of Columns</div>
                    <input
                        type="number"
                        value={cols}
                        onChange={e => setCols(e.target.value)}
                        className="editable-input"
                    />
                    <button
                        className="calculate-button"
                        onClick={handleCalculate}
                        disabled={!rows || !cols}
                    >
                        Calculate
                    </button>
                </div>
            </div>
            {loading ? (
                <div className="loader">Loading...</div>
            ) : (
                <>
                    {matrix.length > 0 && (
                        <div className="table-container">
                            <table className="matrix-table">
                                <tbody>
                                    {matrix.map((row, rowIndex) => (
                                        <tr key={rowIndex}>
                                            {row.map((value, colIndex) => (
                                                <td key={colIndex}>{value}</td>
                                            ))}
                                        </tr>
                                    ))}
                                </tbody>
                            </table>
                        </div>
                    )}
                </>
            )}
            {errorMessage && (
                <div className="error-message">
                    {errorMessage}
                </div>
            )}
        </div>
    );
}

export default Matrix;
