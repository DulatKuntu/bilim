import React from "react";

import "../header-sass/header-main.sass";

import headerMain from "../header-img/img.png";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faSearch } from "@fortawesome/free-solid-svg-icons";

const HeaderMain = () => {
    return (
        <div className="header-main">
            <div className="header-main-text">
                <div className="header-main-text__description">
                    get your
                    <br />
                    <span className="header-main-text__description_header">
                        mentor’s
                    </span>
                    <br />
                    from here
                </div>
                <form className="header-main-text__search">
                    <input
                        type="text"
                        name="search"
                        className="header-main-text__search_input"
                        placeholder="search"
                    />
                    <FontAwesomeIcon
                        icon={faSearch}
                        className="header-main-text__search_emoji"
                    />
                </form>
            </div>

            <div className="header-main-photo-container">
                <img src={headerMain} alt="" className="header-main-photo" />
            </div>
        </div>
    );
};

export default HeaderMain;
