(echo "Initial Problem:")
(echo "")

(declare-const x Int)
(declare-const y Int)
(declare-const z Int)
(declare-const a Int)
(declare-const b Int)

(assert (= a 3 ))
(assert (< 0 x  10))
(assert (< 0 y  10))
(assert (< 0 z  10))
(assert (= (+ a b) z))

(assert (= (+ (* 3 y) (* 2 x)) z))

(check-sat)
(get-model)
(push)

(echo "")
(echo "-----------------------")
(echo "Unsatisfiable Problem:")
(echo "")

(assert (= a b))

(check-sat)
