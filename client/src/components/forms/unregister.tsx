import clsx from "clsx";
import { FunctionComponent } from "react";
import { useForm } from "react-hook-form";
import { ErrorMessage } from "@hookform/error-message";

const Unregister: FunctionComponent = () => {
  const {
    handleSubmit,
    formState: { errors },
    register,
  } = useForm({
    criteriaMode: "all",
  });

  const onSubmit = (data) => console.log(data);

  return (
    <form
      className={clsx("flex", "flex-col")}
      onSubmit={handleSubmit(onSubmit)}
    >
      <p>Enter your phone number to unregister from mfa-spike</p>
      <div className={clsx("mt-2", "flex", "flex-col")}>
        <label>Phone Number</label>
        <input
          {...register("phone", {
            required: "Phone number is required",
            pattern: {
              value: /^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]\d{3}[\s.-]\d{4}$/,
              message: "Invalid phone number",
            },
          })}
          type="tel"
          placeholder="Enter your phone number"
          className={clsx("border", "border-solid", "p-2", "rounded", {
            "border-red-500": errors && errors["phone"],
          })}
        />
        <ErrorMessage
          name="phone"
          errors={errors}
          render={({ messages }) =>
            messages &&
            Object.entries(messages).map(([type, message]) => (
              <p className={clsx("text-red-500", "text-sm")} key={type}>
                {message}
              </p>
            ))
          }
        />
      </div>
      <div className={clsx("mt-2")}>
        <button
          type="submit"
          className={clsx(
            "bg-green-300",
            "pt-2",
            "pb-2",
            "pl-5",
            "pr-5",
            "rounded"
          )}
        >
          Submit
        </button>
      </div>
    </form>
  );
};

export default Unregister;
