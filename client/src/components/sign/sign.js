import React, { useState } from "react";

import SignIn from "./sign-comp/sign-in";
import SignUp from "./sign-comp/sign-up";

import "./sign-sass/sign.sass";

import signMain from "./sign-img/main.png";

const Sign = () => {
    const [sign, setSign] = useState(false);

    return (
        <div className="sign">
            <img src={signMain} alt="" className="sign-main__img" />
            {sign ? <SignUp /> : <SignIn setSign={setSign} />}
        </div>
    );
};

export default Sign;
