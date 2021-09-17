import clsx from "clsx";
import Register from "../components/forms/register";

export default function RegisterPage() {
  return (
    <div>
      <p className={clsx("text-xl", "mt-2", "mb-2")}>Register with mfa-spike</p>
      <div className={clsx("lg:w-1/3", "w-full")}>
        <Register />
      </div>
    </div>
  );
}
