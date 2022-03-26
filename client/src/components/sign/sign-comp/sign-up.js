import React from "react";

import "../sign-sass/sign-up.sass";

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

            <div className="sign-in-text">Sign up!</div>

            <form className="sign-in-main">
                <div className="sign-in-main__inputs">
                    <div className="sign-in-main__inputs_block">
                        <FontAwesomeIcon
                            icon={faEnvelope}
                            className="sign-in-main__inputs_block__emoji"
                            id="sign-in-first__emoji"
                        />

                        <input
                            type="text"
                            name="Email"
                            className="sign-in-main__inputs_block__input"
                            placeholder="example@gmail.com"
                        />

                        <FontAwesomeIcon
                            icon={faKey}
                            className="sign-in-main__inputs_block__emoji"
                            id="sign-in-second__emoji"
                        />

                        <input
                            type="text"
                            name="Password"
                            className="sign-in-main__inputs_block__input"
                            placeholder="Enter your password"
                        />
                    </div>

                    <div className="sign-in-main__inputs_block">
                        <FontAwesomeIcon
                            icon={faKey}
                            className="sign-in-main__inputs_block__emoji"
                            id="sign-in-second__emoji"
                        />

                        <input
                            type="text"
                            name="Password"
                            className="sign-in-main__inputs_block__input"
                            placeholder="Enter your password"
                        />

                        <FontAwesomeIcon
                            icon={faKey}
                            className="sign-in-main__inputs_block__emoji"
                            id="sign-in-second__emoji"
                        />

                        <input
                            type="text"
                            name="Password"
                            className="sign-in-main__inputs_block__input"
                            placeholder="Enter your password"
                        />
                    </div>
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
