import React from "react";

import "../header-sass/header-nav.sass";

import { Link } from "react-router-dom";

const HeaderNav = ({ userRegistered }) => {
    return (
        <div className="header-nav">
            <div className="header-nav-logo__container">
                <img src="" alt="" className="header-nav-logo" />
            </div>

            <div className="header-nav-company">Company name</div>

            <ul className="header-nav-main">
                <li className="header-nav-main__elements">
                    <Link
                        to="/mentors"
                        className="header-nav-main__elements_link"
                    >
                        Mentors
                    </Link>
                </li>
                <li className="header-nav-main__elements">
                    <a href="#" className="header-nav-main__elements_link">
                        Study buddy
                    </a>
                </li>
                <li className="header-nav-main__elements">
                    <a href="#" className="header-nav-main__elements_link">
                        Groups
                    </a>
                </li>
            </ul>

            {userRegistered ? (
                <Link to="/mentors/1" className="header-nav-sign">
                    User Page
                </Link>
            ) : (
                <Link to="/sign" className="header-nav-sign">
                    Login/Sign up
                </Link>
            )}
        </div>
    );
};

export default HeaderNav;
