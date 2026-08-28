import { services } from "./services";

const Header = () => {
    const checkAll = async () => {
    for (const service of services) {
        await checkService(service);
    }
    };

    return (<header className="header">
        <div>
          <h1>Azure Full Stack Health</h1>
          <p>Check the status of every component</p>
        </div>

        <button className="check-all" onClick={checkAll}>
          Check All
        </button>
    </header>)
}

export default Header