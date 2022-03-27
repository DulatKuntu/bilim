import "./survey-sass/survey.sass";

import { Link } from "react-router-dom";

const Survey = () => {
    return (
        <div className="survey">
            <h1 className="survey-header">Опросник</h1>

            <div className="survey-element">
                <h3 className="survey-element__text">
                    1. Are you a numbers whiz?
                </h3>

                <div class="container">
                    <div class="group">
                        <input id="radio-1" name="radios" type="radio" />
                        <label for="radio-1">Yes, I am the next Einstein</label>
                    </div>
                    <div class="group">
                        <input id="radio-2" name="radios" type="radio" />
                        <label for="radio-2">
                            I am good at math but it is not my fav subject
                        </label>
                    </div>
                    <div class="group">
                        <input id="radio-3" name="radios" type="radio" />
                        <label for="radio-3">
                            No, it is really not my thing
                        </label>
                    </div>
                    <div class="group">
                        <input id="radio-4" name="radios" type="radio" />
                        <label for="radio-4">Whats a number</label>
                    </div>
                </div>
            </div>

            <div className="survey-element">
                <h3 className="survey-element__text">
                    2. Do you have a strong creative side?
                </h3>

                <div class="container">
                    <div class="group">
                        <input id="radio-5" name="radios1" type="radio" />
                        <label for="radio-5"> Yes, Picasso, who?</label>
                    </div>
                    <div class="group">
                        <input id="radio-6" name="radios1" type="radio" />
                        <label for="radio-6">
                            I like being creative, but it's more of a hobby
                        </label>
                    </div>
                    <div class="group">
                        <input id="radio-7" name="radios1" type="radio" />
                        <label for="radio-7">No, no, never</label>
                    </div>
                    <div class="group">
                        <input id="radio-8" name="radios1" type="radio" />
                        <label for="radio-8">Yea, Yea</label>
                    </div>
                </div>
            </div>

            <div className="survey-element">
                <h3 className="survey-element__text">
                    3. Are you good at writing
                </h3>

                <div class="container">
                    <div class="group">
                        <input id="radio-9" name="radios2" type="radio" />
                        <label for="radio-9">Yes, I am the next Dickens</label>
                    </div>
                    <div class="group">
                        <input id="radio-10" name="radios2" type="radio" />
                        <label for="radio-10">
                            Yes, just dont ask me to write anything more then
                            few pages
                        </label>
                    </div>
                    <div class="group">
                        <input id="radio-11" name="radios2" type="radio" />
                        <label for="radio-11">
                            I write well but I dont enjoy it so much
                        </label>
                    </div>
                    <div class="group">
                        <input id="radio-12" name="radios2" type="radio" />
                        <label for="radio-12">Nope</label>
                    </div>
                </div>
            </div>
            <Link to="/sign/user/survey/results">
                <button className="survey-go-next">Next</button>
            </Link>
        </div>
    );
};

export default Survey;
