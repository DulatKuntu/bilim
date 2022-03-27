import React, { useState } from "react";

import "./sign-sass/sign-middle.sass";

import { Link } from "react-router-dom";

const SignMiddle = ({ setUser }) => {
    return (
        <div className="sign-middle">
            <Link
                to="/sign/user"
                style={{ textDecoration: "none" }}
                onClick={() => {
                    setUser("s");
                }}
            >
                <div className="sign-middle-container">Я студент!</div>
            </Link>
            <Link
                to="/sign/mentor"
                style={{ textDecoration: "none" }}
                onClick={() => {
                    setUser("m");
                }}
            >
                <div className="sign-middle-container">Я ментор!</div>
            </Link>
        </div>
    );
};

export default SignMiddle;
