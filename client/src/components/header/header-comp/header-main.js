import React from "react";

import "../header-sass/header-main.sass";

import headerMain from "../header-img/bg.jpg";

import { Link } from "react-router-dom";

const HeaderMain = () => {
    return (
        <div className="header-main">
            <div className="header-main-text">
                <div className="header-main-text__description">
                    <div className="header-main-text__description_element">
                        Hold the
                        <span className="header-main-text__description_header">
                            vision
                        </span>
                    </div>
                    <br />
                    <div className="header-main-text__description_element">
                        trust the
                        <span className="header-main-text__description_header">
                            process
                        </span>
                    </div>
                </div>
                <Link
                    to="/sign"
                    className="header-main__search"
                    style={{ textDecoration: "none" }}
                >
                    Искать сейчас
                </Link>
            </div>

            <div className="header-main-photo-container">
                <img src={headerMain} alt="" className="header-main-photo" />
            </div>
        </div>
    );
};

export default HeaderMain;
