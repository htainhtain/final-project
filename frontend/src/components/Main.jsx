import { useState } from "react";
import { services }  from "./services";
import ServiceStatus from "./ServiceStatus";
import { pingpong } from "./pingpong";

const Main = () => {
    const [status, setStatus] = useState({});

    const checkService = async (service) => {
        // Frontend is working if this React application is running.
        if (service.id === "frontend") {
            setStatus((prev) => ({
                ...prev,
                frontend: {
                state: "success",
                message: "Frontend is running",
                },
            }));
        return;
        }

        setStatus((prev) => ({
            ...prev,
            [service.id]: {
                state: "loading",
                message: "Checking...",
            },
        }));

        try {
            const response = await service.endpoint();

            const data = await response.data;

            if (response.status !== 200 ) {
                console.log("data.message : ", data.message )
                throw new Error(data.message || "Request failed");
            }

            setStatus((prev) => ({
                ...prev,
                [service.id]: {
                state: "success",
                message: data.message || "Working",
                },
            }));
        } catch (error) {
            console.log(error)

            setStatus((prev) => ({
                ...prev,
                [service.id]: {
                state: "error",
                message: error.message || "Service unavailable",
                },
            }));
            }
        }
    return (
        <main className="dashboard">
            {services.map((service) => {
                const serviceStatus = status[service.id];

            return (
                <div className="service-card" key={service.id}>
                    <div className="service-info">
                        <h2>{service.name}</h2>
                        <p>{service.description}</p>
                    </div>

                    <ServiceStatus serviceStatus={serviceStatus}/>

                    <button
                        className="check-button"
                        onClick={() => checkService(service)}
                        disabled={serviceStatus?.state === "loading"}
                    >
                        {serviceStatus?.state === "loading"
                        ? "Checking..."
                        : "Check"}
                    </button>
                </div>
                );
            })}
        </main>
    )
}


export default Main