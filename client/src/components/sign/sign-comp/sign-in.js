import React from "react";

import "../sign-sass/sign-in.sass";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faEnvelope, faKey } from "@fortawesome/free-solid-svg-icons";

const SignIn = ({ setSign }) => {
    const handleCard = (event) => {
        event.preventDefault();

        const user = {
            password: event.target.password.value,
            username: event.target.username.value,
        };

        addCardHandler(user);
    };

    async function addCardHandler(card) {
        const response = await fetch("http://localhost:4000/unauthed/signin", {
            method: "POST",
            body: JSON.stringify(card),
            headers: {
                "Content-Type": "application/x-www-form-urlencoded",
                // Accept: "*/*",
            },
        });
        const data = await response.json();
        sessionStorage.setItem("token", JSON.stringify(data));

        getToken();
    }

    function getToken() {
        const tokenString = sessionStorage.getItem("token");
        const userToken = JSON.parse(tokenString);
        return userToken?.token;
    }

    return (
        <div className="sign-in">
            <div className="sign-in-welcome">
                Welcome to
                <br />
                <span className="sign-in-welcome__company">Company name</span>
            </div>

            <div className="sign-in-text">Sign in!</div>

            <form className="sign-in-main" onSubmit={handleCard}>
                <FontAwesomeIcon
                    icon={faEnvelope}
                    className="sign-in-main__emoji"
                    id="sign-in-first__emoji"
                />

                <input
                    type="text"
                    name="username"
                    className="sign-in-main__input"
                    placeholder="example@gmail.com"
                />

                <FontAwesomeIcon
                    icon={faKey}
                    className="sign-in-main__emoji"
                    id="sign-in-second__emoji"
                />

                <input
                    type="text"
                    name="password"
                    className="sign-in-main__input"
                    placeholder="Enter your password"
                />

                <input
                    type="submit"
                    name="Password"
                    className="sign-in-main__submit"
                />
            </form>

            <div className="sign-in-no-account">
                Don't have an account?
                <a
                    href="#"
                    className="sign-in-no-account__link"
                    onClick={() => setSign(true)}
                >
                    Register
                </a>
            </div>
        </div>
    );
};

export default SignIn;
