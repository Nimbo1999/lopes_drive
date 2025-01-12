import { InputHTMLAttributes } from "react";
import clsx from "clsx";

import styles from "./input-group.module.scss";

export interface InputGroupProps extends InputHTMLAttributes<HTMLInputElement> {
  label?: React.ReactNode;
}

export function InputGroup({
  className,
  label,
  id,
  ...props
}: Readonly<InputGroupProps>) {
  return (
    <div className={clsx(styles["input-group"], className)}>
      <label htmlFor={id} className={styles["input-group__label"]}>
        {label}
      </label>

      <input id={id} className={styles["input-group__input"]} {...props} />
    </div>
  );
}
