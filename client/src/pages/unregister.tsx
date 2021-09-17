import clsx from "clsx";
import Unregister from "../components/forms/unregister";

export default function UnregisterPage() {
  return (
    <div>
      <p className={clsx("text-xl", "mt-2", "mb-2")}>
        Unregister from mfa-spike
      </p>
      <div className={clsx("lg:w-1/3", "w-full")}>
        <Unregister />
      </div>
    </div>
  );
}
