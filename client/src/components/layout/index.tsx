import clsx from "clsx";
import { FunctionComponent } from "react";
import Navbar from "./navbar";

const Layout: FunctionComponent = ({ children }) => {
  return (
    <div className={clsx("container", "mx-auto")}>
      <Navbar />
      <div className={clsx("px-4")}>{children}</div>
    </div>
  );
};

export default Layout;
