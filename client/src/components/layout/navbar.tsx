import clsx from "clsx";
import Link from "next/link";

const Navbar = () => {
  return (
    <nav
      className={clsx(
        "flex",
        "bg-gray-100",
        "pt-2",
        "pb-2",
        "pl-1",
        "pr-1",
        "mt-1",
        "rounded"
      )}
    >
      <div className={clsx("flex-grow")}>
        <Link href="/">
          <a className={clsx("hover:underline", "text-blue-500", "ml-2")}>
            mfa-spike
          </a>
        </Link>
      </div>
      <div>
        <Link href="/register">
          <a className={clsx("hover:underline", "text-blue-500", "m-2")}>
            Register
          </a>
        </Link>
        <Link href="/unregister">
          <a className={clsx("hover:underline", "text-blue-500", "m-2")}>
            Unregister
          </a>
        </Link>
        <Link href="/profile">
          <a className={clsx("hover:underline", "text-blue-500", "m-2")}>
            Profile
          </a>
        </Link>
      </div>
    </nav>
  );
};

export default Navbar;
