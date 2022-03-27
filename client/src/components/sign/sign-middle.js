import React, { useState } from "react";

import "./sign-sass/sign-middle.sass";

import { Link } from "react-router-dom";

const SignMiddle = () => {
    return (
        <div className="sign-middle">
            <Link to="/sign/user" style={{ textDecoration: "none" }}>
                <div className="sign-middle-container">Я юзер!</div>
            </Link>
            <Link to="/sign/mentor" style={{ textDecoration: "none" }}>
                <div className="sign-middle-container">Я ментор!</div>
            </Link>
        </div>
    );
};

export default SignMiddle;
