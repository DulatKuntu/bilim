import React from "react";

import HeaderMain from "./header-comp/header-main";
import HeaderNav from "./header-comp/header-nav";

import "./header-sass/header.sass";

const Header = () => {
    return (
        <div className="header">
            <HeaderMain />
        </div>
    );
};

export default Header;
