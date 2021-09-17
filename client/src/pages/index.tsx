import clsx from "clsx";

export default function HomePage() {
  return (
    <div>
      <p className={clsx("text-xl")}>Home</p>
      <p>
        mfa-spike is a sample application intended to explore the implementation
        of multi factor auth. The client is a NextJS application and the server
        is a Golang application. The server utilizes Twilio for SMS
        communication.
      </p>
    </div>
  );
}
