import "../survey-results-sass/survey-results-card.sass";

import { Link } from "react-router-dom";

const SurveyResultsCard = () => {
    return (
        <div className="survey-results-card">
            <div className="survey-results-card-upper">
                <div className="survey-results-card-upper__img_container">
                    <img
                        src=""
                        alt=""
                        className="survey-results-card-upper__img"
                    />
                </div>

                <div className="survey-results-card-upper__text">
                    <div className="survey-results-card-upper__text_main">
                        Test
                    </div>
                    <div className="survey-results-card-upper__text_specialization">
                        Test
                    </div>
                </div>
            </div>

            <div className="survey-results-card-description">
                Lorem ipsum dolor sit amet consectetur adipisicing elit.
                Accusamus, eum maiores ab
            </div>

            <Link
                to="/sign/survey/results/profession"
                className="survey-results-card-link"
            >
                more info
            </Link>
        </div>
    );
};

export default SurveyResultsCard;
