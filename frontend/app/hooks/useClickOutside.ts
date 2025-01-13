import { useEffect } from "react";

/**
 * Hook that alerts clicks outside of the passed ref
 */
export function useClickOutside(
  ref: React.RefObject<HTMLDivElement>,
  callback?: VoidFunction
) {
  useEffect(() => {
    /**
     * Alert if clicked on outside of element
     */
    function handleClickOutside(event: MouseEvent) {
      if (ref.current && !ref.current.contains(event.target as Node)) {
        if (typeof callback === "function") callback();
      }
    }
    // Bind the event listener
    document.addEventListener("click", handleClickOutside);
    return () => {
      // Unbind the event listener on clean up
      document.removeEventListener("click", handleClickOutside);
    };
  }, [ref, callback]);
}
