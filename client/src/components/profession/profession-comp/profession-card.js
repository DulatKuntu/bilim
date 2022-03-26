import "../profession-sass/profession-card.sass";

const ProfessionCard = ({ name, descr }) => {
    return (
        <div className="profession-card">
            <div className="profession-card__text">
                <div className="profession-card__text_main">{name}</div>
            </div>
            <div className="profession-card-description">{descr}</div>
        </div>
    );
};

export default ProfessionCard;
