import React from "react";

import "../sign-sass/sign-in.sass";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faEnvelope, faKey } from "@fortawesome/free-solid-svg-icons";

const SignIn = ({ setSign }) => {
    return (
        <div className="sign-in">
            <div className="sign-in-welcome">
                Welcome to
                <br />
                <span className="sign-in-welcome__company">Company name</span>
            </div>

            <div className="sign-in-text">Sign in!</div>

            <form className="sign-in-main">
                <FontAwesomeIcon
                    icon={faEnvelope}
                    className="sign-in-main__emoji"
                    id="sign-in-first__emoji"
                />

                <input
                    type="text"
                    name="Email"
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
                    name="Password"
                    className="sign-in-main__input"
                    placeholder="Enter your password"
                />

                <div className="sign-in-main__bottom">
                    <div className="sign-in-main__bottom_remember">
                        <input
                            type="checkbox"
                            name="remember"
                            className="sign-in-main__bottom_remember_input"
                        />
                        <div className="sign-in-main__bottom_remember_text">
                            Remember me
                        </div>
                    </div>

                    <a href="#" className="sign-in-main__bottom_forgot">
                        Forgot password
                    </a>
                </div>

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
