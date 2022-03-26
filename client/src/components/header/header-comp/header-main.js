import React from "react";

import "../header-sass/header-main.sass";

import headerMain from "../header-img/bg.jpg";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch } from "@fortawesome/free-solid-svg-icons";

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
                <button className="header-main__search">search now</button>
            </div>

            <div className="header-main-photo-container">
                <img src={headerMain} alt="" className="header-main-photo" />
            </div>
        </div>
    );
};

export default HeaderMain;
