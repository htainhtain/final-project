const ServiceStatus = ({ serviceStatus }) => {
    return (
        <div className="status-container">
            {!serviceStatus && (
                <span className="status unknown">
                    Not checked
                </span>
            )}

            {serviceStatus?.state === "loading" && (
                <span className="status loading">
                    🟡 Checking...
                </span>
            )}

            {serviceStatus?.state === "success" && (
                <span className="status success">
                    🟢 Working
                </span>
            )}

            {serviceStatus?.state === "error" && (
                <span className="status error">
                    🔴 Failed
                </span>
            )}

            {serviceStatus && (
                <small>{serviceStatus.message}</small>
            )}
        </div>
    )
}

export default ServiceStatus