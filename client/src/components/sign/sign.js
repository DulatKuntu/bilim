import React, { useState } from "react";

import SignIn from "./sign-comp/sign-in";
import SignUp from "./sign-comp/sign-up";

import "./sign-sass/sign.sass";

import signMain from "./sign-img/main.png";

const Sign = ({ setUserRegistered, setUserId }) => {
    const [sign, setSign] = useState(false);

    return (
        <div className="sign">
            {sign ? (
                <SignUp setSign={setSign} />
            ) : (
                <SignIn
                    setUserRegistered={setUserRegistered}
                    setSign={setSign}
                    setUserId={setUserId}
                />
            )}
        </div>
    );
};

export default Sign;
